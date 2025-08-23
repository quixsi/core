import { ref } from 'vue'
import { apiService, type User } from '@/services/api'

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
    } catch (e: any) {
        error.value = e.message || 'Failed to create user.'
    }
}

const deleteUser = async (id: string) => {
    error.value = ''
    
    try {
        await apiService.deleteUser(id)
        await fetchUsers()
    } catch (e: any) {
        error.value = e.message || 'Failed to delete user.'
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
