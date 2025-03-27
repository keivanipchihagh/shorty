#!/bin/bash

# Check if the .env file exists
if [[ ! -f .env ]]; then
  echo ".env file not found."
  exit 1
fi

# Create or overwrite the .env.template file
> .env.template

# Read each line of the .env file
while IFS= read -r line || [[ -n "$line" ]]; do
  # Skip empty lines and comments
  if [[ -z "$line" || "$line" =~ ^# ]]; then
    continue
  fi

  key=$(echo "$line" | cut -d'=' -f1)

  # Write key with no value
  echo "$key=" >> .env.template
done < .env

echo ".env.template file created successfully."