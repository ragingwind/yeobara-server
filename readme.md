# Yeobara Check-in Server

> Help attendee to check in their meet-up

## Configuration

You should set up configuration data to `env.json` file before run. Firstly, create a `env.json` file on the root and set configuration data such as app url of firebase. see below sample

```json
{
  "FBAppName": "yeobara.firebaseio.com"
}
```

## Development

### Install bower components

```sh
bower install
```

### Run

```sh
goapp serve
```

### Deploy

```sh
goapp deploy -application [APP-ID]
```

## License

MIT ©[Yeobara](http://github.com/yeobara)
