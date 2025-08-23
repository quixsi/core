import { ref } from 'vue'

export type ToastType = 'success' | 'error'

export interface ToastState {
    message: string
    type: ToastType
    key: number
    duration?: number
    class?: string
}

const toasts = ref<ToastState[]>([])
const timers = new Map<number, ReturnType<typeof setTimeout>>()

export function useToast() {
    const scheduleRemoval = (key: number, duration = 3000) => {
        if (timers.has(key)) return
        const t = setTimeout(() => {
            removeToast(key)
        }, duration)
        timers.set(key, t)
    }

    const showSuccess = (message: string, duration = 3000) => {
        const key = Date.now() + Math.random()
        toasts.value.push({ message, type: 'success', key, duration })
        scheduleRemoval(key, duration)
    }

    const showError = (message: string, duration = 3000) => {
        const key = Date.now() + Math.random()
        toasts.value.push({ message, type: 'error', key, duration })
        scheduleRemoval(key, duration)
    }

    const removeToast = (key: string | number) => {
        toasts.value = toasts.value.filter(t => t.key !== key)
        const timer = timers.get(Number(key))
        if (timer) {
            clearTimeout(timer)
            timers.delete(Number(key))
        }
    }

    const clearToasts = () => {
        toasts.value = []
        timers.forEach(t => clearTimeout(t))
        timers.clear()
    }

    return { toasts, showSuccess, showError, removeToast, clearToasts }
}
