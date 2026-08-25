from ultralytics import YOLO
import cv2
import json

model=YOLO("yolov8n.pt")
image=cv2.imread("road.jpg")

results=model(image)

detections=[]

for result in results:
    for boxes in result.boxes:
        x1,y1,x2,y2=map(int,boxes.xyxy[0])

        confidence=float(boxes.conf[0])
        class_id=int(boxes.cls[0])
        class_name=model.names[class_id]

        detection={
            "class":class_name,
            "confidence":round(confidence,2),
            "bbox":[x1,y1,x2,y2]
        }

        detections.append(detection)

with open("detections.json","w") as file:
    json.dump(detections,file,indent=4)

print("Detection events saved!")