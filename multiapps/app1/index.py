import uvicorn
from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware
from aiofiles import open
from util import Message

app = FastAPI()
app.add_middleware(CORSMiddleware,
                   allow_origins=["*"],
                   allow_credentials=True,
                   allow_methods=["*"],
                   allow_headers=["*"])

# Display message
@app.get("/")
async def fetch_message():
    try:
        async with open("cache.txt", "r") as read_file:
            contents = await read_file.read()
            return {"id": "1", "message": contents}
    except BaseException as error:
        return f"""{error}"""


# Change message
@app.post('/')
async def change_message(*, request: Message):
    try:
        async with open('cache.txt', mode='w') as open_file:
            await open_file.write(request.message)
            return {"id": "1", "message": request.message}
    except BaseException as error:
        return f"""{error}"""


uvicorn.run(app, host='0.0.0.0', port=8080)
