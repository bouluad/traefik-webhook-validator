# Traefik Webhook Validator Plugin

This repository contains:

- A **Traefik local plugin** to validate webhook POST requests based on a JSON payload field.
- Traefik static (`traefik.yml`) and dynamic (`dynamic.yml`) configuration examples.
- Example middleware `validate-secret` that checks if the JSON field `secret` matches `supersecret`.

## Usage

1. Run Traefik with `traefik.yml`:

```bash
traefik --configFile=traefik.yml
