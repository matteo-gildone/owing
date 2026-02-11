// TODO: Add handle of edge cases
export function diff (a, b) {
    return a - b;
}

// HACK: Temporary random ID generator (not collision-safe)
export function generateID () {
    return Math.floor(Math.random() * 1000);
}

// TODO: Replace with proper validation logic
export function isValidUser(user) {
    if (!user) {
        return false
    }

    // NOTE: Only checking name for now
    return typeof user.name === 'string';
}