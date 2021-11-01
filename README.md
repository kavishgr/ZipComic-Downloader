# ZipComic-Downloader

> Wrote this tool to learn Go


## Installation

### From Binary

Download the pre-built binaries for different platforms from the [releases](https://github.com/kavishgr/tempomail/releases/) page. Extract them using tar, move it to your `$PATH` and you're ready to go.

```sh
▶ # download release from https://github.com/kavishgr/?????????/releases/
▶ tar -xzvf linux-amd64-tempomail.tgz
▶ mv tempomail /usr/local/bin/
▶ tempomail -h
```

## Usage

* **Download all comics(directory will be created)**:

```
gozipcomic -u "https://www.zipcomic.com/a-lucky-luke-adventure"
```

* **Specify a range**:

```
gozipcomic -u "https://www.zipcomic.com/a-lucky-luke-adventure" -r '1:10'
```

* **Change directory(should already be available)**

```
gozipcomic -u "https://www.zipcomic.com/a-lucky-luke-adventure" -d '/mydir'
```

* **Help**

```
gozipcomic -h
```
