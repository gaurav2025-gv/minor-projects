from fastapi import FastAPI

app=FastAPI()

items=[]

@app.get("/")
def home():
    return {"message":"FastAPI CRUD API"}

@app.get("/items")
def get_items():
    return items

@app.get("/items/{item_id}")
def get_item(item_id:int):
    for item in items:
        if item["id"]==item_id:
            return item
    return {"error":"Item not found"}

@app.post("/items")
def create_item(item:dict):
    new_item={
        "id":len(items)+1,
        "name":item["name"],
        "severity":item["severity"]
    }

    items.append(new_item)

    return new_item

@app.put("/items/{item_id}")
def update_item(item_id:int,item:dict):
    for existing_item in items:
        if existing_item["id"]==item_id:
            existing_item["name"]=item["name"]
            existing_item["severity"]=item["severity"]
            return existing_item

    return {"error":"Item not found"}

@app.delete("/items/{item_id}")
def delete_item(item_id:int):
    for item in items:
        if item["id"]==item_id:
            items.remove(item)
            return {"message":"Item deleted"}

    return {"error":"Item not found"}