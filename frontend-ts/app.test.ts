import { test, expect } from '@jest/globals';
import { sumar } from './app';

test('suma 2 y 3 para dar 5', () => {
    expect(sumar(2, 3)).toBe(5);
});