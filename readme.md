Steps to repo

1. send a request like below with a breakpoint `if a.CountryCode == "US"` in the resolve func on address
2. Note the first address that comes through works correctly and has "US" assigned
3. Note the second address that resolves has empty string ""

Same issue occurs on the struct passed into the request when looking at i.Body.Away.CountryCode

```
http://localhost:8888/test

{
  "name": "test",
  "age": 6,
  "home": {
    "city": "unknown",
    "line1": "unknown",
    "line2": null,
    "state": "WA",
    "zip": "98001"
  },
  "away": {
    "city": "unknown",
    "line1": "unknown",
    "line2": null,
    "state": "WA",
    "zip": "98001"
  }
}
```

