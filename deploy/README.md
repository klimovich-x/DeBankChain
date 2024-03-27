## Geth init

If you choose to sync from scratch without using our snapshot, you will need to `geth init` the datadir using provided genesis.json.

```
geth init --datadir=$DATADIR /assets/genesis.json
```

### Intiating sequencer node

If you are initiating a sequencer node, then a miner signing key should be imported too.

```
echo "pwd" > $DATADIR/password
echo $SEQUENCER_KEY > $DATADIR/block-signer-key
geth account import --datadir=$DATADIR --password=$DATADIR/password $DATADIR/block-signer-key
```

Then in following op-geth startup command you need to specify:

```
  --mine
  --miner.etherbase=$SEQUENCER_ADDR
  --unlock=$SEQUENCER_ADDR
```