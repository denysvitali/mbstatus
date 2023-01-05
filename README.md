# mbstatus

A small utility that return the DBUS ModemManager status as a JSON 

## Example

```bash
$ ./mbstatus | jq
```

```json
{
  "imei": "000000000000000",
  "operatorName": "Swisscom",
  "operatorCode": "22801",
  "signal": {
    "lte": {
      "rsrp": -114,
      "rsrq": -13,
      "rssi": -81,
      "snr": 0
    }
  }
}
```

## Compiling

### Requirements

- Go 1.19+

### Instructions

```bash
make build
./build/mbstatus
```

## Why?

I use this with my 5G router to monitor its signal strength remotely. 
Together with [openwrt-mobilebroadband-signal](https://github.com/denysvitali/openwrt-mobilebroadband-signal) I am
able to monitor the current status of the LTE / 5G signal from a web page.