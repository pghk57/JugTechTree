# JugTechTree全体像

```mermaid
graph TD;
    subgraph JugTechTree
        技術木
        編集機能
    end

    M[熟練ジャグラー] -- 技術木の構築
        解説の登録
        技術木のコンテンツに関して議論 --- JugTechTree;
    B[初心者] -- 技体系の確認
        解説閲覧 
        技習得状況の記録 --- JugTechTree;
    A[管理者] -- コンテンツの管理
        データのバックアップ
        障害時の復旧
        ログの確認 --- JugTechTree;
```