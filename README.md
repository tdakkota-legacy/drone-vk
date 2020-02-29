# drone-vk 
[![Build Status](https://travis-ci.com/tdakkota/drone-vk.svg?branch=master)](https://travis-ci.com/tdakkota/drone-vk)
[![codecov](https://codecov.io/gh/tdakkota/drone-vk/branch/master/graph/badge.svg)](https://codecov.io/gh/tdakkota/drone-vk)

Drone plugin for sending VK messages


## Usage
```yaml
- name: send vk notification
  image: tdakkota/drone-vk
  settings:
    token: 
      from_secret: token
    peer_id: <user_id> or <chat_id>
```