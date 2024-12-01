package transport

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/adrg/xdg"
	"github.com/metalgrid/drift/internal/platform"
)

func HandleConnection(ctx context.Context, conn net.Conn, gw platform.Gateway) {
	fmt.Println("handling connection", conn.LocalAddr().(*net.TCPAddr), conn.RemoteAddr().(*net.TCPAddr))
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		raw, err := reader.ReadString(byte(endOfMessage))
		if err != nil {
			fmt.Println("error reading from remote:", err)
			return
		}

		msg := UnmarshalMessage(raw)

		switch m := msg.(type) {
		case error:
			gw.Notify(fmt.Sprintf("Error: %s", m))
			return
		case Offer:
			answer := gw.Ask(fmt.Sprintf("Incoming file: %s (%s)", m.Filename, formatSize(m.Size)))
			if answer == "ACCEPT" {
				_, err = conn.Write(Accept().MarshalMessage())
				if err != nil {
					return
				}
			}

			// empty string means waiting for an action from the local user has timed out, so we decline by default
			if answer == "DECLINE" || answer == "" {
				_, _ = conn.Write(Decline().MarshalMessage())
				return
			}

			err = storeFile(xdg.UserDirs.Download+string(os.PathSeparator)+"Drift", m.Filename, m.Size, reader)
			if err != nil {
				gw.Notify(fmt.Sprintf("Failed storing file: %s", err))
				return
			}
			gw.Notify(fmt.Sprintf("File received: %s", m.Filename))
		case Answer:
			if m.Accepted() {
				file := ctx.Value("filename").(string)
				err = sendFile(file, writer)
				if err != nil {
					gw.Notify(fmt.Sprintf("Failed sending %s: %s", file, err))
					return
				}
				gw.Notify(fmt.Sprintf("File sent: %s", file))
			}
		}
	}
}

func storeFile(incoming, file string, size int64, reader *bufio.Reader) error {
	err := os.MkdirAll(incoming, 0777)

	if err != nil {
		return err
	}

	filepath := incoming + string(os.PathSeparator) + file
	// f, err := os.Create(filepath)
	f, err := os.CreateTemp(incoming, file+"*.drift")
	if err != nil {
		return err
	}

	_, err = io.CopyN(f, reader, size)

	_ = f.Close()
	if err != nil {
		// if we're able to create the file, we should be able to remove it as well
		_ = os.Remove(f.Name())
		return err
	}
	return os.Rename(f.Name(), filepath)
}

func sendFile(file string, writer *bufio.Writer) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(writer, f)
	_ = writer.Flush()

	return err
}

func SendFile(filename string, conn net.Conn) {
	fmt.Println("Sending file", filename)
	offer, err := MakeOffer(filename)
	if err != nil {
		fmt.Println("failed creating file offer:", err)
		return
	}

	_, err = conn.Write(offer.MarshalMessage())
	if err != nil {
		fmt.Println("failed sending file:", err)
	}
}
