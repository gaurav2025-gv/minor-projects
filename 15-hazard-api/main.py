from fastapi import FastAPI
import json

app=FastAPI()

FILE="hazards.json"

def load_hazards():
    with open(FILE,"r") as f:
        return json.load(f)

def save_hazards(hazards):
    with open(FILE,"w") as f:
        json.dump(hazards,f,indent=4)

@app.get("/")
def home():
    return {"message":"Hazard API"}

@app.get("/hazards")
def get_hazards():
    return load_hazards()

@app.get("/hazards/{hazard_id}")
def get_hazard(hazard_id:int):
    hazards=load_hazards()

    for hazard in hazards:
        if hazard["id"]==hazard_id:
            return hazard

    return {"error":"Hazard not found"}

@app.post("/hazards")
def create_hazard(hazard:dict):
    hazards=load_hazards()

    new_hazard={
        "id":len(hazards)+1,
        "type":hazard["type"],
        "severity":hazard["severity"],
        "latitude":hazard["latitude"],
        "longitude":hazard["longitude"],
        "confidence":hazard["confidence"]
    }

    hazards.append(new_hazard)
    save_hazards(hazards)

    return new_hazard

@app.delete("/hazards/{hazard_id}")
def delete_hazard(hazard_id:int):
    hazards=load_hazards()

    for hazard in hazards:
        if hazard["id"]==hazard_id:
            hazards.remove(hazard)
            save_hazards(hazards)
            return {"message":"Hazard deleted"}

    return {"error":"Hazard not found"}