import File1 from './File1.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "story/File1",
  component: File1,
  argTypes: { msg: { control: 'text', defaultValue: '' }, bebra: { defaultValue: [] }, bobra: { control: 'number', defaultValue: 0 }, b: { control: 'boolean', defaultValue: false }, heh: { g: { control: 'text', defaultValue: '' } }, beb: { control: 'number', defaultValue: 0 }, },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File1 } from '@estelink/story'\n\n",
          "<File1 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File1 },
    setup: () => ({ args }),
    template: `<File1 v-bind="args" />`,
  });

export const Default = Template.bind({});
