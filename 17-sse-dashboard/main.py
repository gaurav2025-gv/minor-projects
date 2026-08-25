from fastapi import FastAPI
from fastapi.responses import FileResponse, StreamingResponse
import asyncio
import json
import random
from datetime import datetime

app=FastAPI()

@app.get("/")
def home():
    return FileResponse("index.html")

@app.get("/events")
async def events():

    async def generate():

        hazards=["Pothole","Road Crack","Damaged Road"]

        while True:

            hazard={
                "type":random.choice(hazards),
                "severity":random.choice(["Low","Medium","High"]),
                "time":datetime.now().strftime("%H:%M:%S")
            }

            yield f"data: {json.dumps(hazard)}\n\n"

            await asyncio.sleep(2)

    return StreamingResponse(
        generate(),
        media_type="text/event-stream"
    )