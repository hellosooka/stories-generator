import File6 from './File6.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "test/File6",
  component: File6,
  argTypes: {
  },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File6 } from '@estelink/test'\n\n",
          "<File6 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File6 },
    setup: () => ({ args }),
    template: `<File6 v-bind="args" />`,
  });

export const Default = Template.bind({});
