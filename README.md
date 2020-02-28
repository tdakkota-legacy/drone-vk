# drone-vk
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