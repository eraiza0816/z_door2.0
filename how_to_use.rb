# How to use

**このドアは解錠から #{AUTO_LOCK_SEC} 秒後に自動的に鍵が掛かります。**

1. 起動させます。
1. カードを登録します。
1. カードをかざします。
1. 解錠されます。未登録のカードは弾かれます。

カードのログは（RasPiがネットワークに繋がっていれば）Spreadsheetに記録されます。

## 注意事項
- Pasori（リーダー）が外れると停止します。再起動のためにはRasPiのConsoleを開く必要があるので外さないでください。
- ネットワークから切れるとログが記録されません。外さないでください。
  - ネットがなくても開閉は正常に出来ます。
- 未登録のカードを長時間かざすと処理できずに落ちる不具合があります。（修正中）4時間以上かざさないでください。
- なにかあったら[@eraiza0816](https://twitter.com/eraiza0816)まで。

# Log
[hire](https://docs.google.com/spreadsheets/d/16dnM_fcZhj-seAavndLCHZvxERE-6B52Ia5F663mC1Q/edit#gid=0)
