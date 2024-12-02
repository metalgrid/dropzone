// Copyright 2023 Tailscale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package win

const (
	VSCLASS_AEROWIZARDSTYLE      = "AEROWIZARDSTYLE"
	VSCLASS_AEROWIZARD           = "AEROWIZARD"
	VSCLASS_BUTTONSTYLE          = "BUTTONSTYLE"
	VSCLASS_BUTTON               = "BUTTON"
	VSCLASS_COMBOBOXSTYLE        = "COMBOBOXSTYLE"
	VSCLASS_COMBOBOX             = "COMBOBOX"
	VSCLASS_COMMUNICATIONSSTYLE  = "COMMUNICATIONSSTYLE"
	VSCLASS_COMMUNICATIONS       = "COMMUNICATIONS"
	VSCLASS_CONTROLPANELSTYLE    = "CONTROLPANELSTYLE"
	VSCLASS_CONTROLPANEL         = "CONTROLPANEL"
	VSCLASS_DATEPICKERSTYLE      = "DATEPICKERSTYLE"
	VSCLASS_DATEPICKER           = "DATEPICKER"
	VSCLASS_DRAGDROPSTYLE        = "DRAGDROPSTYLE"
	VSCLASS_DRAGDROP             = "DRAGDROP"
	VSCLASS_EDITSTYLE            = "EDITSTYLE"
	VSCLASS_EDIT                 = "EDIT"
	VSCLASS_EXPLORERBARSTYLE     = "EXPLORERBARSTYLE"
	VSCLASS_EXPLORERBAR          = "EXPLORERBAR"
	VSCLASS_FLYOUTSTYLE          = "FLYOUTSTYLE"
	VSCLASS_FLYOUT               = "FLYOUT"
	VSCLASS_HEADERSTYLE          = "HEADERSTYLE"
	VSCLASS_HEADER               = "HEADER"
	VSCLASS_LISTBOXSTYLE         = "LISTBOXSTYLE"
	VSCLASS_LISTBOX              = "LISTBOX"
	VSCLASS_LISTVIEWSTYLE        = "LISTVIEWSTYLE"
	VSCLASS_LISTVIEW             = "LISTVIEW"
	VSCLASS_MENUSTYLE            = "MENUSTYLE"
	VSCLASS_MENU                 = "MENU"
	VSCLASS_NAVIGATION           = "NAVIGATION"
	VSCLASS_PROGRESSSTYLE        = "PROGRESSSTYLE"
	VSCLASS_PROGRESS             = "PROGRESS"
	VSCLASS_REBARSTYLE           = "REBARSTYLE"
	VSCLASS_REBAR                = "REBAR"
	VSCLASS_SCROLLBARSTYLE       = "SCROLLBARSTYLE"
	VSCLASS_SCROLLBAR            = "SCROLLBAR"
	VSCLASS_SPINSTYLE            = "SPINSTYLE"
	VSCLASS_SPIN                 = "SPIN"
	VSCLASS_STATUSSTYLE          = "STATUSSTYLE"
	VSCLASS_STATUS               = "STATUS"
	VSCLASS_TABSTYLE             = "TABSTYLE"
	VSCLASS_TAB                  = "TAB"
	VSCLASS_TASKDIALOGSTYLE      = "TASKDIALOGSTYLE"
	VSCLASS_TASKDIALOG           = "TASKDIALOG"
	VSCLASS_TEXTSTYLE            = "TEXTSTYLE"
	VSCLASS_TOOLBARSTYLE         = "TOOLBARSTYLE"
	VSCLASS_TOOLBAR              = "TOOLBAR"
	VSCLASS_TOOLTIPSTYLE         = "TOOLTIPSTYLE"
	VSCLASS_TOOLTIP              = "TOOLTIP"
	VSCLASS_TRACKBARSTYLE        = "TRACKBARSTYLE"
	VSCLASS_TRACKBAR             = "TRACKBAR"
	VSCLASS_TREEVIEWSTYLE        = "TREEVIEWSTYLE"
	VSCLASS_TREEVIEW             = "TREEVIEW"
	VSCLASS_USERTILE             = "USERTILE"
	VSCLASS_TEXTSELECTIONGRIPPER = "TEXTSELECTIONGRIPPER"
	VSCLASS_WINDOWSTYLE          = "WINDOWSTYLE"
	VSCLASS_WINDOW               = "WINDOW"
)

const (
	MENU_MENUITEM_TMSCHEMA        = 1
	MENU_MENUDROPDOWN_TMSCHEMA    = 2
	MENU_MENUBARITEM_TMSCHEMA     = 3
	MENU_MENUBARDROPDOWN_TMSCHEMA = 4
	MENU_CHEVRON_TMSCHEMA         = 5
	MENU_SEPARATOR_TMSCHEMA       = 6
	MENU_BARBACKGROUND            = 7
	MENU_BARITEM                  = 8
	MENU_POPUPBACKGROUND          = 9
	MENU_POPUPBORDERS             = 10
	MENU_POPUPCHECK               = 11
	MENU_POPUPCHECKBACKGROUND     = 12
	MENU_POPUPGUTTER              = 13
	MENU_POPUPITEM                = 14
	MENU_POPUPSEPARATOR           = 15
	MENU_POPUPSUBMENU             = 16
	MENU_SYSTEMCLOSE              = 17
	MENU_SYSTEMMAXIMIZE           = 18
	MENU_SYSTEMMINIMIZE           = 19
	MENU_SYSTEMRESTORE            = 20
)

const (
	MC_CHECKMARKNORMAL   = 1
	MC_CHECKMARKDISABLED = 2
	MC_BULLETNORMAL      = 3
	MC_BULLETDISABLED    = 4
)

const (
	MCB_DISABLED = 1
	MCB_NORMAL   = 2
	MCB_BITMAP   = 3
)

const (
	MPI_NORMAL      = 1
	MPI_HOT         = 2
	MPI_DISABLED    = 3
	MPI_DISABLEDHOT = 4
)

const (
	MSM_NORMAL   = 1
	MSM_DISABLED = 2
)

const (
	TEXT_MAININSTRUCTION = 1
	TEXT_INSTRUCTION     = 2
	TEXT_BODYTITLE       = 3
	TEXT_BODYTEXT        = 4
	TEXT_SECONDARYTEXT   = 5
	TEXT_HYPERLINKTEXT   = 6
	TEXT_EXPANDED        = 7
	TEXT_LABEL           = 8
	TEXT_CONTROLLABEL    = 9
)

const (
	TS_HYPERLINK_NORMAL   = 1
	TS_HYPERLINK_HOT      = 2
	TS_HYPERLINK_PRESSED  = 3
	TS_HYPERLINK_DISABLED = 4
)

const (
	TS_CONTROLLABEL_NORMAL   = 1
	TS_CONTROLLABEL_DISABLED = 2
)
