## ADDED Requirements

### Requirement: 專案 MUST 提供根目錄 README 總覽
系統 MUST 在倉庫根目錄提供 `README.md`，並以繁體中文說明此專案是一個以 Go 開發、可在 macOS 上執行的 Tetris 遊戲。README MUST 讓第一次進入倉庫的讀者能快速理解專案目的、目前定位與主要技術組成。

#### Scenario: 新讀者查看專案首頁
- **WHEN** 新讀者打開倉庫根目錄並查看 `README.md`
- **THEN** 文件會清楚說明這是一個 Go 實作的 Tetris 專案，並描述其目前聚焦於核心玩法與可維護結構

### Requirement: README MUST 描述目前已支援的核心能力
README MUST 說明目前專案已支援的核心遊戲能力，包括可玩的基本迴圈、核心操作、分數與等級，以及下一個方塊預覽等內容；README MUST 只描述已存在於目前程式或規格中的能力，不可把未實作功能寫成既有能力。

#### Scenario: 讀者確認目前功能範圍
- **WHEN** 讀者閱讀 README 中的功能摘要
- **THEN** 文件只會列出目前已支援或已有規格定義的核心能力，而不會把排行榜、多人模式或其他未實作延伸功能描述為現況

### Requirement: README MUST 提供可執行的開發與驗證入口
README MUST 提供本機開發與驗證入口，至少包含啟動遊戲、執行測試與可用的打包命令或入口說明。這些命令與說明 MUST 對應到倉庫中實際存在的入口，例如 `Makefile` 目標或等價命令。

#### Scenario: 開發者依 README 嘗試執行專案
- **WHEN** 開發者依 README 提供的命令啟動遊戲或執行測試
- **THEN** 所引用的命令與入口存在於目前倉庫中，且不需要額外猜測檔案位置或主要執行方式

### Requirement: README MUST 說明主要模組與責任邊界
README MUST 提供主要目錄或模組的導覽，至少涵蓋 `cmd/tetris`、`internal/app`、`internal/game`、`internal/input` 與 `internal/render` 的責任，讓讀者理解應用入口、遊戲規則、輸入處理與畫面渲染之間的分工。

#### Scenario: 維護者查閱專案結構
- **WHEN** 維護者閱讀 README 的專案結構說明
- **THEN** 維護者可以辨識各主要模組的責任邊界，而不需要先逐一打開所有原始碼檔案

### Requirement: README MUST 明確界定目前限制與文件可信來源
README MUST 說明專案目前的限制或範圍界線，並以現有程式碼、`Makefile` 與 OpenSpec 規格作為描述依據。當 README 提及操作方式、功能摘要或結構資訊時，內容 MUST 能被目前倉庫中的對應實作或規格交叉驗證。

#### Scenario: 維護者檢查 README 是否可信
- **WHEN** 維護者拿 README 與現有命令、入口檔案及規格比對
- **THEN** README 中的功能、命令與結構描述可以在目前倉庫中找到對應來源，且不包含無法驗證的宣稱
