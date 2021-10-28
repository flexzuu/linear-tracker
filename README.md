![Screenshot 2021-04-06 at 15 55 47](https://user-images.githubusercontent.com/3532919/113722371-a638f480-96f0-11eb-8443-73efa55c4dfc.png)


# install
```bash
go install ./cmd/linear-tracker
```

# run

```bash
linear-tracker -token=<token>
```

# how to create a token
visit https://linear.app/settings/api and create a personal api key

# how to keep alive

- add a lauchd plist into your `~/Library/LaunchAgents/` folder
- remember to replace the path, working dir, fullpath to binary and token