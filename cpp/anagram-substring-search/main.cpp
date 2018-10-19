// Problem Description:
// https://www.geeksforgeeks.org/anagram-substring-search-search-permutations/

#include <iostream>
#include <unordered_set>

using namespace std;

void initCharArray(int* arr) {
    for (int i = 0; i < 256; i++) {
        arr[i] = 0;
    }
}

bool set_contains(unordered_set<char>& hashSet, char ch) {
    return hashSet.find(ch) != hashSet.end();
}

int anagramsCount(string pattern, string text) {
    int totalOccurences = 0;
    int charSequenceCnt = 0;
    int charF[256] = { 0 };
    int expectedCharF[256] = { 0 };
    int expectedCharSeqCnt = 0;
    unordered_set<char> patternLetters;

    for (int i = 0; i < pattern.length(); i++) {
        patternLetters.insert(pattern[i]);

        if (expectedCharF[pattern[i]] == 0) {
            expectedCharSeqCnt++;
        }

        expectedCharF[pattern[i]]++;
    }

    for (int i = 0; i < text.length(); i++) {
        if (i > pattern.length() - 1) {
            char letterToDrop = text[i - pattern.length()];
            if (set_contains(patternLetters, letterToDrop)) {

                int previousCharF = charF[letterToDrop];
                charF[letterToDrop]--;

                if (previousCharF == expectedCharF[letterToDrop]) {
                    charSequenceCnt--;
                }
            }
        }

        char letterToAdd = text[i];
        if (set_contains(patternLetters, letterToAdd)) {

            charF[letterToAdd]++;
            if (charF[letterToAdd] == expectedCharF[letterToAdd]) {
                charSequenceCnt++;
            }
        }

        if (charSequenceCnt == expectedCharSeqCnt) {
            totalOccurences++;
        }
    }

    return totalOccurences;
}

int main() {
	int testCases = 0;
	string text;
	string pattern;

	cin >> testCases;
	for (int i = 0; i < testCases; i++) {
	    cin >> text;
	    cin >> pattern;

	    cout << anagramsCount(pattern, text) << endl;
	}

	return 0;
}
