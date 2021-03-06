# SSPro-Server

## Intro
Monitoring focusing on collecting sensors data from RPI and sending to remote node.
Data should be collected and sent to remote server where it will be presented via WebUI.
Alerting for critical tresholds should be used.

## Note
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

- Correct build script to create main file to run in scratch
- Source gosrc to create alias 'gobld main .'
- main = output file 

## Input
Datapoints needs to have following format:
`JSON:`
  {'data':[datapoint1, datapoint2, ...] }

datapointX structure:
```
  { 'timestamp':<int>,
    'location':<string>,
    'room':<string>,
    'name':<string>,
    'sensor':<string>,
    'value':<real/float>
  }
```

Explanation:
 - location = provides multi tenancy (multisite) deployment
 - room = place inside the location
 - name = sensor name, just for redundancy if we add 2 or more same sensor type device
 - sensor = sensor type like: humidity, temperature, fire, smoke, ...
 - value = value generated by sensor or recalculated

### API

- POST - /login - Generate token
- POST - /datapoints - Save new datapoint
- GET  - /datapoints/all - Show datapoints all
- GET  - /datapoints/<location> - show all datapoints for locatoin only
- GET  - /datapoints/<location>/<room> - show all datapoits per room
- GET  - /datapoints/<location>/<room>/<name> - show all datapoints per sensor name


### MVP
- API to store datapoints per location and time
- HTTPS
- Authentication
- Small DB
- Show all datapoints via get as JSON (no GUI)
