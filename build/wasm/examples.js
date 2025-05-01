// Store the last selected example
let lastSelectedExample = '';

// Load example code from file
async function loadExampleCode(filename) {
    try {
        const response = await fetch(`examples/${filename}.vx`);
        if (!response.ok) throw new Error(`Failed to load ${filename}`);
        return await response.text();
    } catch (error) {
        console.error('Error loading example:', error);
        return null;
    }
}

async function loadExample() {
    const select = document.getElementById('codeExamples');
    const selectedExample = select.value;
    
    if (selectedExample && selectedExample !== '') {
        const code = await loadExampleCode(selectedExample);
        if (code) {
            window.editor.setValue(code);
            lastSelectedExample = selectedExample;
        }
    } else if (selectedExample === '') {
        // If user selects the default option, restore the last selected example
        if (lastSelectedExample) {
            select.value = lastSelectedExample;
            const code = await loadExampleCode(lastSelectedExample);
            if (code) {
                window.editor.setValue(code);
            }
        }
    }
} 