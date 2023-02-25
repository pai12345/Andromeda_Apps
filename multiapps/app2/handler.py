from aiohttp import ClientSession
import json

# Class for processing message
class Message():
    # Method for fetching message
    async def fetch(self, request):
        try:
            async with ClientSession() as session:
                async with session.get(request) as resp:
                    result = await resp.json()
                    print(resp)
                    return result
        except BaseException as error:
            return f"""{error}"""
    # Method for reversing message
    def reverse(self, request):
        try:
            result = json.loads(json.dumps(request))
            result["message"] = result["message"][::-1]
            return result
        except BaseException as error:
            return f"""{error}"""


message = Message()
