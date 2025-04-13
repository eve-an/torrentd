# Torrentd

Ever downloaded a torrent over a torrent client like [qBittorrent](https://www.qbittorrent.org/) or [Transmission](https://transmissionbt.com/) and you wanted to automate the completion of the torrent?
Often these clients don't handle finished torrents well (see [qbit Issue](https://github.com/qbittorrent/qBittorrent/issues/21568)).
To be independent of the torrent client behaviour and its quirks **Torrentd** comes into place!

**Torrentd** monitors the status of a torrent download. Currently it is designed to do so as a cron or CLI, but in future versions a daemon should run in background
to fulfill additional tasks like moving completed files to a given destination. 

## Usage

```shell
./torrentd <file> <torrentFile>
```

## Technical background

### Bittorrent protocol

The protocol can be found [here](https://www.bittorrent.org/beps/bep_0003.html).
For us the important sections are *[bencoding](https://en.wikipedia.org/wiki/Bencode)*, *metainfo files* and *trackers*.

When we want to download a torrent we first need to download the corresponding `.torrent` file.
This is a *bencoded* file with meta information for our torrent client.
The bencoded file saves the URL of the tracker alongside with an `info` node.
There we find information like the name of the downloaded file, it's length and most importantly:
`pieces` and `pieces length`.

`pieces length` gives us the chunk size in bytes for the hashing of the original file. That means our 
file is chunked and for each chunk we compute the *[sha1](https://en.wikipedia.org/wiki/SHA-1)* hash for progress monitoring.
The counterpart for that is the `pieces` field where the hashes are stored as byte strings.

A bencoded `archlinux-2025.03.01-x86_64.iso.torrent` file json transformed (https://archlinux.org/download/).
```json
{
  "comment": "Arch Linux 2025.03.01 <https://archlinux.org>",
  "created by": "mktorrent 1.1",
  "creation date": 1740850982,
  "info": {
    "length": 1247838208,
    "name": "archlinux-2025.03.01-x86_64.iso",
    "piece length": 524288,
    "pieces": "..."
  },
  "url-list":["..."]
}
```

### Bencoding Grammar

```
<value>     ::= <integer> | <string> | <list> | <dictionary>

<integer>   ::= 'i' <digits> 'e'
<string>    ::= <length>:<data>
<list>      ::= 'l' <value>* 'e'
<dictionary>::= 'd' (<string> <value>)* 'e'
```

## TODOs

- [X] migrate to Cobra from custom args parsing
- [ ] move from Cron like system to a daemon
- [ ] add extensible system for executing commands after torrent completion
- [ ] add support for multi file downloads
