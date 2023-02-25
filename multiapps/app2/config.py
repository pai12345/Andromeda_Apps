import os

# Fetch environment variables
def get_env():
    return {
        "TARGET_HOST": os.environ['TARGET_HOST'],
        "TARGET_PORT": os.environ['TARGET_PORT']
    }
