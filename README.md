# klaytn-geth-db-analysis

> This repository for explaining Kalytn db data

Dependency
- https://github.com/klaytn/klaytn
- https://github.com/ian0371/homi-local

## Output

```shell
├── db_extract.txt # DB extract
├── db_extract.md # descrition
└── db
    ├── 0_genesis_setup
    ├── 1_mine_1_block
    └── 4_mine_4_block
```
> [db_extract.md](./db_extract.md)

## Extract DB

```
go run main.go > db_extract.txt
```

## How to Run

#### 1. Clone Klaytn

```shell
$ git clone https://github.com/klaytn/klaytn.git
```

#### 2. Build

```shell
/klaytn $ make all
```

#### 3. Clone homi-local

```shell
/klaytn $ git clone https://github.com/ian0371/homi-local
```
 
#### 4. Setup Path
```shell
/klaytn $ export PATH=$PATH:~/klaytn/build/bin

/klaytn $ which homi
~/klaytn/build/bin/homi
```

#### 5. Edit Shard Number


```diff
-     kcn --datadir "$DATA_DIR" init "$DATA_DIR/genesis.json"
+     kcn --db.num-statetrie-shards 1 --datadir "$DATA_DIR" init "$DATA_DIR/genesis.json"
```
> [homi-local/cmd/setup.sh#L11](https://github.com/ian0371/homi-local/blob/main/cmd/setup.sh#L11)


```diff
-    $BIN/kcn $OPTIONS >> ${LOG_DIR}/kcnd.out 2>&1 &
+    $BIN/kcn --db.num-statetrie-shards 1 $OPTIONS >> ${LOG_DIR}/kcnd.out 2>&1 &
```
> [homi-local/kcnd.sh#L287](https://github.com/ian0371/homi-local/blob/main/kcnd#L287)

#### 6. Setup Genesis Chaindata
```shell
/klaytn/homi-local $ homi setup --cn-num 1 --funding-addr ${address} local
/klaytn/homi-local $ ./run setup
```
> More details from [homi-local](https://github.com/ian0371/homi-local)

#### 7. Run Consensus Node

```shell
/klaytn/homi-local $ ./run start
/klaytn/homi-local $ ./run stop
```
