$(function() {
    $('#analyseBtn').click(function() {
        $.ajax({
            url: '/analyse',
            data: JSON.stringify({ 'inputText': $('#inputText').val() }),
            type: 'POST',
            contentType: 'application/json',
            success: function(response) {
                handleResponse(response);
                encodeCSV(response);
            },
            error: function(error) {
                console.log(error);
            }
        });
    });
});

function handleResponse(response) {
    $('#lexicon_count').text("" + response.lexicon_count);
    $('#syllable_count').text("" + response.syllable_count);
    $('#sentence_count').text("" + response.sentence_count);
    $('#flesch_reading_ease').text("" + response.flesch_reading_ease);
    $('#flesch_kincaid_grade').text("" + response.flesch_kincaid_grade);
    $('#gunning_fog').text("" + response.gunning_fog);
    $('#smog_index').text("" + response.smog_index);
    $('#automated_readability_index').text("" + response.automated_readability_index);
    $('#coleman_liau_index').text("" + response.coleman_liau_index);
    $('#linsear_write_formula').text("" + response.linsear_write_formula);
    $('#dale_chall_readability_score').text("" + response.dale_chall_readability_score);
}

var csvRows = [];

function encodeCSV(response) {
    csvRows = [];
    csvRows.push(['']);
    csvRows.push(['lexicon_count', "" + response.lexicon_count]);
    csvRows.push(['syllable_count', "" + response.syllable_count]);
    csvRows.push(['sentence_count', "" + response.sentence_count]);
    csvRows.push(['flesch_reading_ease', "" + response.flesch_reading_ease]);
    csvRows.push(['flesch_kincaid_grade', "" + response.flesch_kincaid_grade]);
    csvRows.push(['gunning_fog', "" + response.gunning_fog]);
    csvRows.push(['smog_index', "" + response.smog_index]);
    csvRows.push(['automated_readability_index', "" + response.automated_readability_index]);
    csvRows.push(['coleman_liau_index', "" + response.coleman_liau_index]);
    csvRows.push(['linsear_write_formula', "" + response.linsear_write_formula]);
    csvRows.push(['dale_chall_readability_score', "" + response.dale_chall_readability_score]);
}

$(function() {
    $('#exportToCSV').click(function() {
        var csvContent = "data:text/csv;charset=utf-8,";
        csvRows.forEach(function(currentRowArray) {
            var row = currentRowArray.join(',');
            csvContent += row + "\r\n";
        });

        // Append hidden link that downloads the CSV file and click it
        var encodedUri = encodeURI(csvContent);
        var link = document.createElement("a");
        link.setAttribute("href", encodedUri);
        link.setAttribute("download", "my_data.csv");

        document.body.appendChild(link);
        link.click();
    });
});
