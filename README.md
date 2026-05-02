# BMSDB

BMSDB is a golang library helps reading local data from bms clients' local data and scan files under a possible client installation. The supported clients are `LR2` and `Beatoraja`. See test files to know how to use this library.

## Supported tables

### Beatoraja

| Table | Supported? | Api | Explain |
| ------------- | -------------- | -------------- | -------------- |
| songdata.db/folder | No | - | - |
| songdata.db/song | Yes | SongData | - |
| songinfo.db/information | Yes | SongInfo | - |
| score.db/score | Yes | Score | - |
| score.db/player | No | - | - |
| score.db/info | No | - | - |
| scorelog.db/scorelog | Yes | ScoreLog | - |
| scoredatalog.db/scoredatalog | Yes | ScoreDataLog | - |

### LR2

| Table | Supported | Api | Explain |
| ------------- | -------------- | -------------- | -------------- |
| \[User.db\]/score | Yes | Score | - |
| \[User.db\]/player | No | - | - |
| song.db/song | Yes |  Song | - |
| song.db/expert | No | - | - |
| song.db/folder | No | - | - |
| song.db/grade | No | - | - |
| song.db/nonstop | No | - | - |

## Tests

Tests are based on real database files. To run the test you have to open a '.database' directory under the project. And place the database files under it. The directory structure follows the client installation.

Example:

```
.
├── beatoraja
│   ├── player
│   │   ├── player1
│   │   │   ├── score.db
│   │   │   ├── scoredatalog.db
│   │   │   └── scorelog.db
│   │   └── player2
│   │       ├── score.db
│   │       ├── scoredatalog.db
│   │       └── scorelog.db
│   ├── songdata.db
│   ├── songinfo.db
│   ├── songinfo.db-shm
│   └── songinfo.db-wal
└── LR2
    └── LR2files
        └── Database
            ├── Score
            │   └── ctz.db
            └── song.db
```

## LICENSE

MIT
