import pandas as pd
import json

with open("./users.json", "r") as f:
    users = json.load(f)
    for i, user in enumerate(users):
        user["email"] = f"mail{i + 1}@maill.com"
        user["password"] = "12345678"
        users[i] = user
    users = pd.DataFrame(users)
    users = users[["id", "name", "email", "password", "avatar", "created_at"]]
    users.to_csv("./users.csv", index=False)
