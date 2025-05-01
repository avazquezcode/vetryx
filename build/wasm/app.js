// Initialize Monaco Editor
require.config({ paths: { 'vs': 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.45.0/min/vs' }});
require(['vs/editor/editor.main'], function() {
    // Create editor instance
    const editor = monaco.editor.create(document.getElementById('editor'), {
        value: '## Enter your Vetryx code here\n',
        language: 'python', // We'll use python highlighting for now (since it's closer in syntax).
        theme: 'vs-dark',
        automaticLayout: true,
        minimap: {
            enabled: false
        },
        fontSize: 14,
        lineNumbers: 'on',
        roundedSelection: false,
        scrollBeyondLastLine: false,
        readOnly: false,
        cursorStyle: 'line',
        tabSize: 4,
        insertSpaces: true
    });

    // Make editor available globally
    window.editor = editor;
});

// Initialize the WASM module
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

// Add keyboard shortcut (Ctrl+Enter or Cmd+Enter)
document.addEventListener('keydown', function(e) {
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        runCode();
    }
});
