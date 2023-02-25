import uvicorn
from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware
from handler import message
from config import get_env

app = FastAPI()
app.add_middleware(CORSMiddleware,
                   allow_origins=["*"],
                   allow_credentials=True,
                   allow_methods=["*"],
                   allow_headers=["*"])

# API call to fetch message
@app.get("/")
async def default():
    try:
        TARGET_HOST = get_env()["TARGET_HOST"]
        TARGET_PORT = get_env()["TARGET_PORT"]
        fetch_message = await message.fetch(f"http://{TARGET_HOST}:{TARGET_PORT}")
        reverse_message = message.reverse(fetch_message)
        return reverse_message
    except BaseException as error:
        return f"""{error}"""


uvicorn.run(app, host='0.0.0.0', port=8081)
