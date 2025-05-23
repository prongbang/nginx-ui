<script setup lang="ts">
import type { SelectProps } from 'ant-design-vue'

const props = defineProps<{
  mask?: Record<string | number, string | (() => string)> | (() => Promise<Record<string | number, string>>)
  placeholder?: string
  multiple?: boolean
  // eslint-disable-next-line ts/no-explicit-any
  defaultValue?: any
}>()

const selectedValue = defineModel<string | number | string[] | number[]>('value')
const options = ref<SelectProps['options']>([])

async function loadOptions() {
  options.value = []
  let actualValue: number | string
  if (typeof props.mask === 'function') {
    const getOptions = props.mask as (() => Promise<Record<string | number, string>>)

    const r = await getOptions()
    for (const [value, label] of Object.entries(r)) {
      actualValue = value
      if (typeof selectedValue.value === 'number')
        actualValue = Number(value)
      options.value?.push({ label, value: actualValue })
    }

    return
  }
  for (const [value, label] of Object.entries(props.mask as Record<string | number, string | (() => string)>)) {
    let actualLabel = label

    if (typeof label === 'function')
      actualLabel = label()

    actualValue = value
    if (typeof selectedValue.value === 'number')
      actualValue = Number(value)

    options.value?.push({ label: actualLabel, value: actualValue })
    if (actualValue === selectedValue.value)
      selectedValue.value = actualValue
  }
}

function init() {
  loadOptions()
}

watch(props, init)

onMounted(() => {
  if (!selectedValue.value && props.defaultValue)
    selectedValue.value = props.defaultValue

  init()
})
</script>

<template>
  <ASelect
    v-model:value="selectedValue"
    allow-clear
    :options="options"
    :placeholder="props.placeholder"
    :default-active-first-option="false"
    :mode="props.multiple ? 'multiple' : undefined"
    class="min-w-180px w-auto!"
    :get-popup-container="triggerNode => triggerNode.parentNode"
  />
</template>

<style lang="less" scoped></style>
