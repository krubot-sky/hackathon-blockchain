echo 'Configtxgen Starts';
ls -l /shared;

sleep 1 && while [ ! -f /shared/status_cryptogen_complete ];
do
  echo Waiting for cryptogen;
  sleep 1;
done;

cp /shared/artifacts/configtx.yaml /shared/;
cd /shared/;
export FABRIC_CFG_PATH=$PWD;

configtxgen -profile TwoOrgsOrdererGenesis -outputBlock genesis.block && find /shared -type d | xargs chmod a+rx && find /shared -type f | xargs chmod a+r && touch /shared/status_configtxgen_complete && rm /shared/status_cryptogen_complete;
