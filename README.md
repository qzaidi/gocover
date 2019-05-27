
The uniq identifier to use for your repo is 

```
git config --get remote.origin.url | md5
```

To get the coverage badge, add this to readme

```
http://localhost:8080/img/cover/<unique id>.png

```

To push coverage data from your CI

```
curl -X PUT -d <coverage> http://zaidi.ml/cover/<unique id>

```
