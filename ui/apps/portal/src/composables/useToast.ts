import { ref } from 'vue'

export type ToastType = 'success' | 'error'

interface ToastState {
    message: string
    type: ToastType
    key: number
}

const toast = ref<ToastState>({ message: '', type: 'success', key: 0 })

export function useToast() {
    const showSuccess = (message: string) => {
        toast.value = { message, type: 'success', key: Date.now() }
    }

    const showError = (message: string) => {
        toast.value = { message, type: 'error', key: Date.now() }
    }

    const clearToast = () => {
        toast.value.message = ''
    }
  
    return { toast, showSuccess, showError, clearToast }
}
