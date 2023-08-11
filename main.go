package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	TRUE  = 1
	FALSE = 0
	NULL  = 0

	SPI_GETMENUANIMATION          = 0x1002 // 11? 將功能表淡出或滑動到檢視
	SPI_SETMENUANIMATION          = 0x1003 // 11? 將功能表淡出或滑動到檢視
	SPI_GETCOMBOBOXANIMATION      = 0x1004 // 13. 滑動開啟下拉試方塊
	SPI_SETCOMBOBOXANIMATION      = 0x1005 // 13. 滑動開啟下拉試方塊
	SPI_GETLISTBOXSMOOTHSCROLLING = 0x1006 // 03. 平滑捲動的清單方塊
	SPI_SETLISTBOXSMOOTHSCROLLING = 0x1007 // 03. 平滑捲動的清單方塊
	SPI_GETMENUFADE               = 0x1012 // 11. 將功能表淡出或滑動到檢視
	SPI_SETMENUFADE               = 0x1013 // 11. 將功能表淡出或滑動到檢視
	SPI_GETSELECTIONFADE          = 0x1014 // 08. 按下功能表項目後淡出
	SPI_SETSELECTIONFADE          = 0x1015 // 08. 按下功能表項目後淡出
	SPI_GETTOOLTIPANIMATION       = 0x1016 // 10? 將工具提示淡出或滑動到檢視
	SPI_SETTOOLTIPANIMATION       = 0x1017 // 10? 將工具提示淡出或滑動到檢視
	SPI_GETTOOLTIPFADE            = 0x1018 // 10. 將工具提示淡出或滑動到檢視
	SPI_SETTOOLTIPFADE            = 0x1019 // 10. 將工具提示淡出或滑動到檢視
	SPI_GETCURSORSHADOW           = 0x101A // 06. 在滑鼠指標下顯示陰影
	SPI_SETCURSORSHADOW           = 0x101B // 06. 在滑鼠指標下顯示陰影
	SPI_GETDROPSHADOW             = 0x1024 // 04. 在視窗下顯示陰影
	SPI_SETDROPSHADOW             = 0x1025 // 04. 在視窗下顯示陰影
	SPI_SETDRAGFULLWINDOWS        = 0x0025 // 07. 拖曳時顯示視窗內容
	SPI_GETDRAGFULLWINDOWS        = 0x0026 // 07. 拖曳時顯示視窗內容
	SPI_GETUIEFFECTS              = 0x103E // XX.
	SPI_SETUIEFFECTS              = 0x103F // XX.
	SPI_GETCLIENTAREAANIMATION    = 0x1042 // 05. 在視窗內部以動畫方式顯示控制項和元素
	SPI_SETCLIENTAREAANIMATION    = 0x1043 // 05. 在視窗內部以動畫方式顯示控制項和元素
	SPI_GETANIMATION              = 0x0048 // 12. 將視窗最大化或最小化時顯示視窗動畫
	SPI_SETANIMATION              = 0x0049 // 12. 將視窗最大化或最小化時顯示視窗動畫
	SPI_GETFONTSMOOTHING          = 0x004A // 02. 去除螢幕字型毛邊
	SPI_SETFONTSMOOTHING          = 0x004B // 02. 去除螢幕字型毛邊

	SPIF_UPDATEINIFILE = 0x01
	SPIF_SENDCHANGE    = 0x02
)

var (
	user32Handle             windows.Handle
	systemParametersInfoProc uintptr
)

func init() {
	if h, err := windows.LoadLibrary("User32.dll"); err != nil {
		panic(err)
	} else if p, err := windows.GetProcAddress(h, "SystemParametersInfoA"); err != nil {
		panic(err)
	} else {
		user32Handle = h
		systemParametersInfoProc = p
	}
}

func main() {
	defer windows.FreeLibrary(user32Handle)

	// 01. 工具列中的動畫
	// TODO

	// 02. 去除螢幕字型毛邊
	fmt.Print("02. FONTSMOOTHING: ")
	if enabled, err := getSystemParameterInfo(SPI_GETFONTSMOOTHING); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETFONTSMOOTHING, true, false); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 03. 平滑捲動的清單方塊
	fmt.Print("03. LISTBOXSMOOTHSCROLLING: ")
	if enabled, err := getSystemParameterInfo(SPI_GETLISTBOXSMOOTHSCROLLING); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETLISTBOXSMOOTHSCROLLING, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 04. 在視窗下顯示陰影
	fmt.Print("04. DROPSHADOW: ")
	if enabled, err := getSystemParameterInfo(SPI_GETDROPSHADOW); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETDROPSHADOW, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 05. 在視窗內部以動畫方式顯示控制項和元素
	fmt.Print("05. CLIENTAREAANIMATION: ")
	if enabled, err := getSystemParameterInfo(SPI_GETCLIENTAREAANIMATION); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETCLIENTAREAANIMATION, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 06. 在滑鼠指標下顯示陰影
	fmt.Print("06. CURSORSHADOW: ")
	if enabled, err := getSystemParameterInfo(SPI_GETCURSORSHADOW); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETCURSORSHADOW, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 07. 拖曳時顯示視窗內容
	fmt.Print("07. DRAGFULLWINDOWS: ")
	if enabled, err := getSystemParameterInfo(SPI_GETDRAGFULLWINDOWS); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETDRAGFULLWINDOWS, true, false); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 08. 按下功能表項目後淡出
	fmt.Print("08. MENUANIMATION: ")
	if enabled, err := getSystemParameterInfo(SPI_GETMENUANIMATION); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETMENUANIMATION, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 09. 啟用 Peek

	// 10. 將工具提示淡出或滑動到檢視
	fmt.Print("10. TOOLTIPANIMATION: ")
	if enabled, err := getSystemParameterInfo(SPI_GETTOOLTIPANIMATION); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETTOOLTIPANIMATION, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}
	fmt.Print("10. TOOLTIPFADE: ")
	if enabled, err := getSystemParameterInfo(SPI_GETTOOLTIPFADE); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETTOOLTIPFADE, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 11. 將功能表淡出或滑動到檢視
	fmt.Print("11. MENUANIMATION: ")
	if enabled, err := getSystemParameterInfo(SPI_GETMENUANIMATION); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETMENUANIMATION, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}
	fmt.Print("11. MENUFADE: ")
	if enabled, err := getSystemParameterInfo(SPI_GETMENUFADE); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETMENUFADE, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// 12. 將視窗最大化或最小化時顯示視窗動畫
	// TODO

	// 13. 滑動開啟下拉試方塊
	fmt.Print("13. COMBOBOXANIMATION: ")
	if enabled, err := getSystemParameterInfo(SPI_GETCOMBOBOXANIMATION); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETCOMBOBOXANIMATION, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}

	// XX.
	fmt.Print("XX. UIEFFECTS: ")
	if enabled, err := getSystemParameterInfo(SPI_GETUIEFFECTS); err != nil {
		panic(err)
	} else if !enabled {
		fmt.Println("enabling...")
		if err := setSystemParameterInfo(SPI_SETUIEFFECTS, true, true); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("already enabled")
	}
}

func getSystemParameterInfo(action uint) (bool, error) {
	enabled := false
	if ok, _, err := syscall.SyscallN(
		uintptr(systemParametersInfoProc),
		uintptr(action),
		NULL,
		uintptr(unsafe.Pointer(&enabled)),
		NULL,
	); ok != TRUE {
		return false, err
	}
	return enabled, nil
}

func setSystemParameterInfo(action uint, enabled bool, pv bool) error {
	e := FALSE
	if enabled {
		e = TRUE
	}
	enabledUI := NULL
	enabledPV := NULL
	if pv {
		enabledPV = e
	} else {
		enabledUI = e
	}
	if ok, _, err := syscall.SyscallN(
		uintptr(systemParametersInfoProc),
		uintptr(action),
		uintptr(enabledUI),
		uintptr(enabledPV),
		uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE),
	); ok != TRUE {
		return err
	}
	return nil
}
