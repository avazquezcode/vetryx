// Initialize Monaco Editor
require.config({ paths: { 'vs': 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.45.0/min/vs' }});
require(['vs/editor/editor.main'], function () {
    initVetryxLanguage();

    const editor = monaco.editor.create(document.getElementById('editor'), {
        value: '',
        language: 'vetryx',
        theme: 'vs-dark',
        automaticLayout: true,
        minimap: { enabled: false },
        fontSize: 14,
        lineNumbers: 'on',
        roundedSelection: false,
        scrollBeyondLastLine: false,
        readOnly: false,
        cursorStyle: 'line',
        tabSize: 4,
        insertSpaces: true,
        wordWrap: 'on',
        wrappingStrategy: 'advanced',
        fixedOverflowWidgets: true
    });

    window.addEventListener('resize', () => {
        editor.layout();
    });

    window.editor = editor;

    fetch('examples/helloworld.vx')
        .then(response => response.text())
        .then(code => {
            editor.setValue(code);
            setTimeout(() => editor.layout(), 0);
        })
        .catch(error => {
            console.error('Error loading hello world example:', error);
        });
});

// Initialize WASM
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);
    });

function runCode() {
    const code = window.editor.getValue();
    const output = document.getElementById('output');

    try {
        const result = window.compileAndRun(code);
        if (result.error) {
            output.innerHTML = `<span class="error">Error: ${result.error}</span>`;
        } else {
            output.textContent = result.output;
        }
    } catch (error) {
        output.innerHTML = `<span class="error">Error: ${error.message}</span>`;
    }
}

document.addEventListener('keydown', function (e) {
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        runCode();
    }
});
