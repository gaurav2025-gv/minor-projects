from flask import Flask,request,jsonify,render_template
import json

app=Flask(__name__)

@app.route("/")
def home():
    return render_template("index.html")

@app.route("/location",methods=["POST"])
def location():
    data=request.json

    latitude=data["latitude"]
    longitude=data["longitude"]

    hazard={
        "class":"pothole",
        "confidence":0.91,
        "bbox":[120,80,350,280],
        "latitude":latitude,
        "longitude":longitude
    }

    with open("hazard.json","w") as file:
        json.dump(hazard,file,indent=4)

    print("Latitude:",latitude)
    print("Longitude:",longitude)

    return jsonify({"message":"Location saved!"})

app.run(debug=True)