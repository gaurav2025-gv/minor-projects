from fastapi import FastAPI
from fastapi.responses import FileResponse

app=FastAPI()

@app.get("/")
def home():
    return {"message":"Video API Server"}

@app.get("/video")
def get_video():
    return FileResponse(
        "videos/road.mp4",
        media_type="video/mp4"
    )