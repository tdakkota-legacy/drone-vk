# drone-vk 
[![Build Status](https://travis-ci.com/tdakkota/drone-vk.svg?branch=master)](https://travis-ci.com/tdakkota/drone-vk)
[![docker build](https://img.shields.io/docker/automated/tdakkota/drone-vk)](https://hub.docker.com/r/tdakkota/drone-vk)
[![Go Report Card](https://goreportcard.com/badge/github.com/tdakkota/drone-vk)](https://goreportcard.com/report/github.com/tdakkota/drone-vk)
[![Docker Pulls](https://img.shields.io/docker/pulls/tdakkota/drone-vk.svg)](https://hub.docker.com/r/tdakkota/drone-vk/)
[![microbadger](https://images.microbadger.com/badges/image/tdakkota/drone-vk.svg)](https://microbadger.com/images/tdakkota/drone-vk "Get your own image badge on microbadger.com")

Drone plugin for sending VK messages. Built with [vksdk].


## Usage
```yaml
- name: send vk notification
  image: tdakkota/drone-vk
  settings:
    token: 
      from_secret: token
    peer_id: <user_id> or <chat_id>
```

## License

[BSD-3-Clause](LICENSE)

[vksdk]: https://github.com/SevereCloud/vksdk