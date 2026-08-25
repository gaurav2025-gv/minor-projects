from ultralytics import YOLO
import cv2

model=YOLO("yolov8n.pt")

image=cv2.imread("road.jpg")

results=model(image)

for result in results:
    boxes=result.boxes

    for box in boxes:
        x1,y1,x2,y2=map(int,box.xyxy[0].cpu().numpy())

        confidence=float(box.conf[0].cpu().numpy())
        class_id=int(box.cls[0].cpu().numpy())

        class_name=model.names[class_id]

        width=x2-x1
        height=y2-y1
        area=width*height

        print("Object:",class_name)
        print("Width:",width)
        print("Height:",height)
        print("Area:",area)

        cv2.rectangle(
            image,
            (x1,y1),
            (x2,y2),
            (0,255,0),
            2
        )

        cv2.putText(
            image,
            f"{class_name} {confidence:.2f}",
            (x1,y1-10),
            cv2.FONT_HERSHEY_SIMPLEX,
            0.6,
            (0,255,0),
            2
        )

cv2.imwrite("analyzed.jpg",image)

print("Analysis complete!")