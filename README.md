# lmz

Fetch and change the power status of a La Marzocco Linea, and possibly other La Marzocco machines.

## Usage

```bash
% cp pkg/config/config.yaml.example pkg/config/config.yaml
% vim pkg/config/config.yaml
% make build
% ./lmz 
Status as of 2024-04-13 12:33:37.677 -0700 PDT: StandBy
% ./lmz on
OK
% ./lmz
Status at 2024-04-13 14:16:49.192 -0700 PDT: BrewingMode
```

## See Also

- [This Thread](https://community.home-assistant.io/t/la-marzocco-gs-3-linea-mini-support/203581/2)
- [This Code](https://github.com/rccoleman/lmdirect/blob/master/lmdirect/connection.py)
