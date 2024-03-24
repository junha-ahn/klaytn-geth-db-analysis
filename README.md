# klaytn-geth-db-analysis

> This repository for explaning klaytn/kalytn data structure

Dependency
- https://github.com/klaytn/klaytn
- https://github.com/ian0371/homi-local

## Output Directory

- 0_genesis_setup
-

## Extract DB

```
go run main.go > db_extract.txt
```

## How to Run

### Setup Consensus Node

##### 1. Clone Klaytn

```shell
$ git clone https://github.com/klaytn/klaytn.git
```

##### 2. Build

```shell
/klaytn $ make all
```

##### 3. Clone homi-local

```shell
/klaytn $ git clone https://github.com/ian0371/homi-local
```
 
##### 4. Setup Path
```shell
/klaytn $ export PATH=$PATH:~/project/klaytn/build/bin

/klaytn $ which homi
~/project/klaytn/build/bin/homi
```

##### 5. Edit Shard Number


```diff
-     kcn --datadir "$DATA_DIR" init "$DATA_DIR/genesis.json"
+     kcn --db.num-statetrie-shards 1 --datadir "$DATA_DIR" init "$DATA_DIR/genesis.json"
```
> https://github.com/ian0371/homi-local/blob/main/cmd/setup.sh#L11


```diff
-    $BIN/kcn $OPTIONS >> ${LOG_DIR}/kcnd.out 2>&1 &
+    $BIN/kcn --db.num-statetrie-shards 1 $OPTIONS >> ${LOG_DIR}/kcnd.out 2>&1 &
```
> https://github.com/ian0371/homi-local/blob/main/kcnd#L287

> homi-local/kcnd.sh
##### 6. Run Setup Genesis
```shell
/klaytn/homi-local $ homi setup --cn-num 1 local
/klaytn/homi-local $ ./run setup
```
> More details from [homi-local](https://github.com/ian0371/homi-local)

### Run Consensus Node

```shell
/klaytn/homi-local $ ./run start
/klaytn/homi-local $ ./run stop
```
