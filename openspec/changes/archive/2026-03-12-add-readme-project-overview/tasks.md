## 1. 盤點 README 所需資訊

- [x] 1.1 盤點目前可由程式碼、`Makefile` 與 OpenSpec 規格確認的專案定位、功能摘要與開發入口
- [x] 1.2 整理 `cmd/tetris`、`internal/app`、`internal/game`、`internal/input`、`internal/render` 的責任邊界，作為 README 的結構導覽內容

## 2. 建立 README 內容

- [x] 2.1 在倉庫根目錄新增 `README.md`，以繁體中文撰寫專案介紹、目前功能範圍與技術背景
- [x] 2.2 在 README 中加入執行遊戲、執行測試與 macOS 打包相關命令或入口說明
- [x] 2.3 在 README 中加入專案結構導覽與目前限制，避免描述未實作能力

## 3. 驗證文件一致性

- [x] 3.1 逐項比對 README 與 `Makefile`、程式入口及 `openspec/specs/playable-core-loop/spec.md`，修正文案不一致之處
- [x] 3.2 手動檢查 README 是否涵蓋專案定位、功能摘要、開發命令、模組分工與限制說明
