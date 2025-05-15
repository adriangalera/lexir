import json
import urllib.parse

filename = "../kaikki-spanish.jsonl"
line_limit = 5
count = 0
with open(filename, encoding="utf-8") as f:
    for line in f:

        print(line)
        """
        data = json.loads(line)
        word = data["word"]
        pos_title = data.get("pos_title", "N/A")
        senses = data["senses"]
        sense_length = len(senses)
        sanitized_word = urllib.parse.quote(word)
        link = f"https://es.wiktionary.org/wiki/{sanitized_word}"
        print(f"{word}: {pos_title} {sense_length} significados. Link: {link}")
        """

        count = count + 1
        if count == line_limit:
            exit(0)
