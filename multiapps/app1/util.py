from pydantic import BaseModel


class Message(BaseModel):
    """Class for defining payload type"""
    message: str
