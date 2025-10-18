import pandas as pd
import json

with open("./comments.json", "r") as f:
    comments = json.load(f)
    comments = pd.DataFrame(comments)
    comments["parent_id"] = comments["parent_id"].astype("Int64")

    comments.to_csv("./comments.csv", index=False)
