import { ref } from 'vue'
import { apiService, type User } from '@/services/api'
import { useToast } from '@/composables/useToast'

const { showSuccess, showError } = useToast()
const users = ref<User[]>([])
const loading = ref(true)
const error = ref('')

const fetchUsers = async () => {
    loading.value = true
    error.value = ''

    try {
        users.value = await apiService.getUsers()
    } catch (e: any) {
        error.value = e.message || 'Failed to load users.'
        showError(error.value)
    } finally {
        loading.value = false
    }
}

const createUser = async (name: string, email: string) => {
    error.value = ''
    try {
        if (!name.trim()) {
            throw new Error('Name is required.')
        }
        if (!email.trim()) {
            throw new Error('Email is required.')
        }
        
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
        if (!emailRegex.test(email)) {
            throw new Error('Invalid email address.')
        }

        if (users.value.some(u => u.email.toLowerCase() === email.trim().toLowerCase())) {
            throw new Error('Email is already in use.')
        }
        
        await apiService.createUser({ name, email })
        await fetchUsers()
        showSuccess('User created successfully!')
    } catch (e: any) {
        error.value = e.message || 'Failed to create user.'
        showError(error.value)
    }
}

const deleteUser = async (id: string) => {
    error.value = ''

    try {
        await apiService.deleteUser(id)
        await fetchUsers()
        showSuccess('User deleted successfully!')
    } catch (e: any) {
        error.value = e.message || 'Failed to delete user.'
        showError(error.value)
    }
}

export function useUsers() {
    return {
        users,
        loading,
        error,
        fetchUsers,
        createUser,
        deleteUser,
    }
}
