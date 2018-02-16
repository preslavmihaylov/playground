from flask import Flask, render_template, request, jsonify
from textstat.textstat import textstat

app = Flask(__name__)

@app.route('/')
def main():
    return render_template('index.html')

@app.route('/analyse', methods=['POST'])
def analyseText():
    values = request.get_json()
    required = [ 'inputText' ]
    if not all(k in values for k in required):
        return 'Missing values', 400

    text = values['inputText']
    result = {
        'syllable_count': textstat.syllable_count(text),
        'lexicon_count': textstat.lexicon_count(text),
        'sentence_count': textstat.sentence_count(text),
        'flesch_reading_ease': textstat.flesch_reading_ease(text),
        'flesch_kincaid_grade': textstat.flesch_kincaid_grade(text),
        'gunning_fog': textstat.gunning_fog(text),
        'smog_index': textstat.smog_index(text),
        'automated_readability_index': textstat.automated_readability_index(text),
        'coleman_liau_index': textstat.coleman_liau_index(text),
        'linsear_write_formula': textstat.linsear_write_formula(text),
        'dale_chall_readability_score': textstat.dale_chall_readability_score(text)
    };

    return jsonify(result), 200

if __name__ == "__main__":
    app.run()
