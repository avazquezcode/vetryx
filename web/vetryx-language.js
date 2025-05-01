// Vetryx language definition for Monaco Editor
function initVetryxLanguage() {
    monaco.languages.register({ id: 'vetryx' });

    // Define the language configuration
    monaco.languages.setMonarchTokensProvider('vetryx', {
        // Keywords
        keywords: [
            'fn', 'return', 'if', 'else', 'while', 'dec', 'print', 'break', 'continue', 'true', 'false', 'null'
        ],

        // Built-in functions
        builtinFunctions: [
            'min', 'max', 'sleep', 'clock'
        ],

        // Operators
        operators: [
            '=', ':=', '+', '-', '*', '%', '/', '==', '<>', '<', '>', '<=', '>=', "!", "&&", "||"
        ],

        // Symbols
        symbols: /[=:=+\-*/(){}[\],;!]/,

        // Tokenizer rules
        tokenizer: {
            root: [
                // Comments
                [/##.*$/, 'comment'],
                [/#.*$/, 'comment'],

                // Keywords
                [/\b(fn|return|if|else|while|dec|print|break|continue|true|false|null)\b/, 'keyword'],

                // Built-in functions
                [/\b(min|max|sleep|clock)\b/, 'function'],

                // Operators
                [/[=:=+\-*/%&&||]/, 'operator'],
                [/[=!<>]=?/, 'operator'],

                // Numbers
                [/\d+/, 'number'],

                // Strings
                [/"([^"\\]|\\.)*$/, 'string.invalid'],
                [/'([^'\\]|\\.)*$/, 'string.invalid'],
                [/"/, 'string', '@string_double'],
                [/'/, 'string', '@string_single'],

                // Identifiers
                [/[a-zA-Z_]\w*/, 'identifier'],

                // Whitespace
                [/[ \t\r\n]+/, 'white'],
            ],

            string_double: [
                [/[^\\"]+/, 'string'],
                [/\\./, 'string.escape'],
                [/"/, 'string', '@pop']
            ],

            string_single: [
                [/[^\\']+/, 'string'],
                [/\\./, 'string.escape'],
                [/'/, 'string', '@pop']
            ]
        }
    });

    // Define the language configuration
    monaco.languages.setLanguageConfiguration('vetryx', {
        brackets: [
            ['{', '}'],
            ['(', ')']
        ],
        autoClosingPairs: [
            { open: '{', close: '}' },
            { open: '(', close: ')' },
            { open: '"', close: '"' },
            { open: "'", close: "'" }
        ],
        surroundingPairs: [
            { open: '{', close: '}' },
            { open: '(', close: ')' },
            { open: '"', close: '"' },
            { open: "'", close: "'" }
        ]
    });
} 