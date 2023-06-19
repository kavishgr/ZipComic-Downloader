# ZipComic-Downloader

> For Learning Purposes. Use at your own risk!


## Installation

### From Binary

Download the pre-built binaries for different platforms from the [releases](https://github.com/kavishgr/ZipComic-Downloader/releases) page. Extract them using tar, move it to your `$PATH` and you're ready to go.

> Did not test on Windows

```sh
▶ # download release from https://github.com/kavishgr/https://github.com/kavishgr/ZipComic-Downloader/releases/
▶ tar -xzvf linux-amd64-gozipcomic.tgz
▶ mv gozipcomic /usr/local/bin/
▶ gozipcomic -h
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

* **Change directory(directory should already be available)**

```
gozipcomic -u "https://www.zipcomic.com/a-lucky-luke-adventure" -d '/mydir'
```

* **Help**

```
gozipcomic -h
```
