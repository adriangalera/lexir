import json

filename = "./kaikki-spanish.jsonl"
line_limit = 100
count = 0
with open(filename, encoding="utf-8") as f:
    for line in f:
        data = json.loads(line)
        word = data["word"]
        pos_title = data.get("pos_title", None)
        senses = data["senses"]
        sense_length = len(senses)
        print(f"{word}: {pos_title}.{sense_length}")
        count = count + 1
        if count == line_limit:
            exit(0)
