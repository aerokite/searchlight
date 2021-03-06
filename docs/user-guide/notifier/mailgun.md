### Notifier `mailgun`

This will send a notification email via mailgun.

#### Configure

To set `mailgun` as notifier, we need to set following environment variables in Icinga2 deployment.

```yaml
env:
  - name: NOTIFY_VIA
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: notify_via
  - name: MAILGUN_DOMAIN
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: mailgun_domain
  - name: MAILGUN_API_KEY
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: mailgun_api_key
  - name: MAILGUN_PUBLIC_API_KEY
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: mailgun_public_api_key
  - name: MAILGUN_FROM
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: mailgun_from
  - name: MAILGUN_TO
    valueFrom:
      secretKeyRef:
        name: appscode-icinga
        key: mailgun_to
```

##### envconfig for `mailgun`

| Name                    | Description                                                                    |
| :---                    | :---                                                                           |
| MAILGUN_DOMAIN          | Set domain name for mailgun configuration                                      |
| MAILGUN_API_KEY         | Set mailgun API Key                                                            |
| MAILGUN_PUBLIC_API_KEY  | Set mailgun public API Key                                                     |
| MAILGUN_FROM            | Set sender address for notification                                            |
| MAILGUN_TO              | Set recipient address. For multiple receipents, set comma separated addresses. |


These environment variables will be set using `appscode-icinga` Secret.

> Set `NOTIFY_VIA` to `mailgun`

#### Set Environment Variables

##### Key `notify_via`
Encode and set `NOTIFY_VIA` to it
```sh
export NOTIFY_VIA=$(echo "mailgun" | base64  -w 0)
```

##### Key `mailgun_domain`
Encode and set `MAILGUN_DOMAIN` to it
```sh
export MAILGUN_DOMAIN=$(echo <domainn> | base64  -w 0)
```

##### Key `mailgun_api_key`
Encode and set `MAILGUN_API_KEY` to it
```sh
export MAILGUN_API_KEY=$(echo <api key> | base64  -w 0)
```

##### Key `mailgun_public_api_key`
Encode and set `MAILGUN_PUBLIC_API_KEY` to it
```sh
export MAILGUN_PUBLIC_API_KEY=$(echo <public api key> | base64  -w 0)
```

##### Key `mailgun_from`
Encode and set `MAILGUN_FROM` to it
```sh
export MAILGUN_FROM=$(echo <sender email address> | base64  -w 0)
```

##### Key `mailgun_to`
Encode and set `MAILGUN_TO` to it
```sh
export MAILGUN_TO=$(echo <recipient email addresses> | base64  -w 0)
```
