// Click handler for the history button
document.addEventListener('DOMContentLoaded', function() {
    const historyBtn = document.querySelector('.history-btn');
    
    historyBtn.addEventListener('click', function() {
        this.classList.toggle('active');
        const historyContent = document.getElementById('historyContent');
        if (historyContent) {
            if (!this.classList.contains('active')) {
                historyContent.classList.add('hidden');
            }
        }
    });
});

// Copy to clipboard function with temporary notification
function copyToClipboard(text) {
    navigator.clipboard.writeText(window.location.origin + '/' + text)
        .then(() => {
            // Create notification element
            const notification = document.createElement('div');
            notification.className = 'copy-notification';
            notification.textContent = 'URL copied to clipboard!';
            document.body.appendChild(notification);

            // Show notification with animation
            setTimeout(() => {
                notification.classList.add('show');
            }, 10);

            // Remove notification after 4 seconds
            setTimeout(() => {
                notification.classList.remove('show');
                setTimeout(() => {
                    notification.remove();
                }, 300); // Wait for fade out animation
            }, 3000);
        })
        .catch(err => {
            console.error('Failed to copy:', err);
            alert('Failed to copy URL');
        });
}